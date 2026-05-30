package errors

type Code string

const (
	CodeInvalidInput         Code = "INVALID_INPUT"
	CodeUnauthorized         Code = "UNAUTHORIZED"
	CodeForbidden            Code = "FORBIDDEN"
	CodeNotFound             Code = "NOT_FOUND"
	CodeConflict             Code = "CONFLICT"
	CodeDuplicateTransaction Code = "DUPLICATE_TRANSACTION"
	CodeInsufficientPoints   Code = "INSUFFICIENT_POINTS"
	CodeRewardExpired        Code = "REWARD_EXPIRED"
	CodeConsentRequired      Code = "CONSENT_REQUIRED"
	CodeInternal             Code = "INTERNAL_ERROR"
)
