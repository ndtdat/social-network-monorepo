# Agents of the system

My system has 2 agents: Operator and User.

Operator has roles (not yet support):
- Create campaigns, configure vouchers, and define relationships between them.
- Stop an active campaign or voucher when needed.
- View statistics.

Users are able to:
- Register with an optional campaign code.
- Purchase a subscription plan to upgrade to a higher tier if desired.

---
# Services and their description

My system includes two services:
- User service:
    - Role: Manages users and campaigns.
    - APIs:
        - Register (public).
        - Login (public).
    - Cronjobs:
        - Monitor campaigns: Checks and marks campaigns as unavailable if they are expired.
- Purchase service:
    - Role: Manages voucher codes, subscription plans, and user tiers.
    - APIs:
        - BuySubscriptionPlan (requires access token).
    - Cronjobs:
        - Monitor voucher configuration: Checks and marks voucher configurations as unavailable if they are expired.
        - Monitor user voucher: Checks and marks user vouchers as unavailable if they are expired.
        - Provision voucher codes for pool: Generates a sufficient number of voucher codes for the pool.

---
# Service explanation
My system includes two services:
- User service:
    - Role: Manages users and campaigns.
    - Description:
        - Register (API):
            - (1) Check if the account already exists and return a response if it does.
            - (2) Check if the campaign code is valid. If invalid, proceed to (4).
            - (3) Trigger the allocation of a voucher for the user.
            - (4) Create the user account and return an access token.
            - Note: The operations of creating a user, updating the campaign, and allocating a voucher must be performed as a transaction to ensure a rollback in case of an error.
        - Login (API):
            - (1) Check if the account exists and return a failure response if not.
            - (2) Verify if the password is correct and return a failure response if incorrect.
            - (3) Return an access token upon successful authentication.
        - Monitor campaign (task background):
            - Runs every minute via a cron job.
            - Marks campaigns as "Unavailable" if end_at is less than the current time.
- Purchase service:
    - Role: Manages subscription plans and vouchers.
    - Description:
        - Buy Subscription Plan (API):
            - (1) Retrieve the user's balance from the payment service (coming soon).
            - (2) Fetch the subscription plan information and the current user's tier.
            - (3) Check if the tier the user wants to purchase is higher than their current tier.
            - (4) Verify if the user has any valid voucher to apply to the purchase. If a valid voucher exists, calculate the discount amount.
            - (5) Perform the following operations in a transactional manner:
              - Mark the user voucher as used.
              - Update the user's tier.
              - Create a transaction record with the purchase details.
              - Call the payment service to debit the user's balance (coming soon).
        - Allocate Voucher (Internal call):
            - (1) Check for a valid voucher based on the campaign ID. The voucher must be in "available" status and not expired.
            - (2) Create a voucher for the user in a transactional operation:
              - Retrieve a voucher code from the pool.
              - Create a user voucher with "Allocated" status.
              - Update the voucher configuration to increase allocated_qty. If allocated_qty reaches max_qty, mark the configuration as "Unavailable".
              - Remove the used voucher code from the pool.
        - Monitor voucher configuration (background task):
            - Runs every minute via a cron job.
            - Marks voucher configurations as "Unavailable" if end_at is less than the current time.
        - Monitor user voucher (background task):
            - Runs every minute via a cron job.
            - Marks user vouchers as "Expired" if expire_at is less than the current time.
        - Provision voucher code pool (background task):
            - Runs every 5 minutes via a cron job.
            - Generates enough voucher codes to maintain the target pool quantity:
              - (1) Generate missing quantities according to rules.
              - (2) Ensure generated codes do not already exist in the pool.
              - (3) Verify that generated codes have not been used by any user.
              - (4) Repeat step (1) if the pool does not have enough codes.
              - (5) Add generated codes to the pool.