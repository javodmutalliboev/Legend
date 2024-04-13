package interface_package

type Order interface {
	Create() (*int64, error)
	GetGoods() error
	ToggleDelivered() error
	ToggleCanceled() error
}
