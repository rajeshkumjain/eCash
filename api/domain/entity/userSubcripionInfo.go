package entity

// combination of all of this is what user subscription would be
// UserSubscriptionInfo: Structure for active user subscription Information
type UserSubscriptionInfo struct {
	Customer     RegisteredUser
	Plan         SubscriptionPlans
	Subscription UserSubscription
}
