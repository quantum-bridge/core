package handlers

import (
	"github.com/quantum-bridge/core/cmd/service/api/requests"
	"github.com/quantum-bridge/core/cmd/service/api/responses"
	"math"
	"net/http"
	"strconv"
)

// GetHistory handles the history endpoint request
// @Summary Get transaction history
// @Description Get combined history of deposits and withdrawals with filtering and pagination
// @ID getHistory
// @Tags History
// @Accept json
// @Produce json
// @Param from_address query string false "Filter by from address"
// @Param to_address query string false "Filter by to address"
// @Param source_network query string false "Filter by source network"
// @Param destination_network query string false "Filter by destination network"
// @Param token_address query string false "Filter by token address"
// @Param transaction_type query string false "Filter by transaction type (deposits/withdrawals)"
// @Param from_block query integer false "Filter by minimum block number"
// @Param to_block query integer false "Filter by maximum block number"
// @Param page query integer false "Page number (default: 1)"
// @Param page_size query integer false "Page size (default: 10)"
// @Param sort_by query string false "Sort field (default: block_number)"
// @Param sort_order query string false "Sort order (asc/desc, default: desc)"
// @Success 200 {object} responses.HistoryResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /history [get]
func GetHistory(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	request, err := requests.NewHistoryRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get repositories
	depositsRepo := DepositsHistoryQuery(r.Context())
	withdrawalsRepo := WithdrawalsHistoryQuery(r.Context())

	// Initialize result slice with estimated capacity
	result := make([]responses.HistoryEntry, 0, request.PageSize)
	var totalCount int64

	// Get deposits if needed
	if request.TransactionType == nil || *request.TransactionType == "deposits" {
		// Apply filters
		query := depositsRepo.New()

		if request.FromAddress != nil {
			query = query.Where("from_address", *request.FromAddress)
		}
		if request.ToAddress != nil {
			query = query.Where("to_address", *request.ToAddress)
		}
		if request.SourceNetwork != nil {
			query = query.Where("source_network", *request.SourceNetwork)
		}
		if request.DestinationNetwork != nil {
			query = query.Where("destination_network", *request.DestinationNetwork)
		}
		if request.TokenAddress != nil {
			query = query.Where("token_address", *request.TokenAddress)
		}
		if request.FromBlock != nil {
			query = query.Where("block_number >=", strconv.FormatUint(*request.FromBlock, 10))
		}
		if request.ToBlock != nil {
			query = query.Where("block_number <=", strconv.FormatUint(*request.ToBlock, 10))
		}

		// Apply sorting
		query = query.OrderBy(request.SortBy, request.SortOrder)

		// Apply pagination
		query = query.Limit(uint64(request.PageSize))

		// Get deposits
		deposits, err := query.Select()
		if err != nil {
			Log(r.Context()).Errorf("failed to get deposits: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert to history entries
		for _, d := range deposits {
			result = append(result, responses.HistoryEntry{
				Type:               "deposit",
				TxHash:             d.TxHash,
				BlockNumber:        d.BlockNumber,
				TokenAddress:       d.TokenAddress,
				TokenID:            d.TokenID,
				Amount:             d.Amount,
				FromAddress:        d.FromAddress,
				ToAddress:          d.ToAddress,
				SourceNetwork:      d.SourceNetwork,
				DestinationNetwork: d.DestinationNetwork,
				IsMintable:         d.IsMintable,
			})
		}

		totalCount += int64(len(deposits))
	}

	// Get withdrawals if needed
	if request.TransactionType == nil || *request.TransactionType == "withdrawals" {
		// Apply filters
		query := withdrawalsRepo.New()

		if request.ToAddress != nil {
			query = query.Where("to_address", *request.ToAddress)
		}
		if request.DestinationNetwork != nil {
			query = query.Where("destination_network", *request.DestinationNetwork)
		}
		if request.TokenAddress != nil {
			query = query.Where("token_address", *request.TokenAddress)
		}
		if request.FromBlock != nil {
			query = query.Where("block_number >=", strconv.FormatUint(*request.FromBlock, 10))
		}
		if request.ToBlock != nil {
			query = query.Where("block_number <=", strconv.FormatUint(*request.ToBlock, 10))
		}

		// Apply sorting
		query = query.OrderBy(request.SortBy, request.SortOrder)

		// Apply pagination
		query = query.Limit(uint64(request.PageSize))

		// Get withdrawals
		withdrawals, err := query.Select()
		if err != nil {
			Log(r.Context()).Errorf("failed to get withdrawals: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert to history entries
		for _, w := range withdrawals {
			result = append(result, responses.HistoryEntry{
				Type:               "withdrawal",
				BlockNumber:        w.BlockNumber,
				TokenAddress:       w.TokenAddress,
				TokenID:            w.TokenID,
				Amount:             w.Amount,
				ToAddress:          w.ToAddress,
				DestinationNetwork: w.DestinationNetwork,
				IsMintable:         w.IsMintable,
				WithdrawalTxHash:   w.WithdrawalTxHash,
				DepositTxHash:      w.DepositTxHash,
			})
		}

		totalCount += int64(len(withdrawals))
	}

	// Prepare response
	response := responses.HistoryResponse{
		Data: result,
		Pagination: struct {
			CurrentPage int   `json:"current_page"`
			PageSize    int   `json:"page_size"`
			TotalItems  int64 `json:"total_items"`
			TotalPages  int   `json:"total_pages"`
		}{
			CurrentPage: request.Page,
			PageSize:    request.PageSize,
			TotalItems:  totalCount,
			TotalPages:  int(math.Ceil(float64(totalCount) / float64(request.PageSize))),
		},
	}

	respond(w, response)
}
