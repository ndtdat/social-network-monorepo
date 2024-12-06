# Agents of system

My system has 2 agents: Operator and User.

Operator has roles (not support yet):
- Create campaigns, voucher configuration and relation between them.
- Can stop a running campaign or a running voucher as them need.
- Can statistics.

User able to:
- Register with campaign code (optional).
- Buy subscription plan if want to upgrade higher tier.

---
# Services and its description

My system includes 2 services:
- User service:
    - Role: Manage users and campaigns.
    - APIs:
        - Register (public).
        - Login (public).
    - Cronjobs:
        - Monitor campaigns: Check and mark unavailable if campaign is expired.
- Purchase service:
    - Role: Manage voucher codes, subscription plan and user tier.
    - APIs:
        - BuySubscriptionPlan (required access token).
    - Cronjobs:
        - Monitor voucher configuration: Check and mark unavailable if voucher configuration is expired.
        - Monitor user voucher: Check and mark unavailable if user voucher is expired.
        - Provision voucher codes for pool: Generate enough codes for pool.

---
# Service explanation
My system includes 2 services:
- User service:
    - Role: Manage users and campaigns.
    - Description:
        - Register (API):
            - (1) Check if account is existed => Return if existed.
            - (2) Check if campaign code is valid => Ignore if invalid and go to (4)
            - (3) Trigger allocate voucher for this user.
            - (4) Create users and return access token.
            - Note: Create user, update campaign and allocate voucher must be in transaction operation because they are rollback when error.
        - Login (API):
            - (1) Check if account is existed => Return fail if not existed.
            - (2) Check if password is correct => Return fail if incorrect.
            - (3) Return access token.
        - Monitor campaign (task background):
            - Run per 1 minutes by cronjob.
            - Mark Unavailable for campaign has end_at < now.
- Purchase service:
    - Role: Manage subscription plan and voucher.
    - Description:
        - Buy Subscription Plan (API):
            - (1) Get user balance from payment-service (coming soon).
            - (2) Get subscription plan info and current user's tier.
            - (3) Check if tier that user want to buy is greater than current tier.
            - (4) Check if user has any valid voucher to apply for this buying => If exists, calculate discount amount.
            - (5) In transactional operation:
                - Mark user voucher is used.
                - Update new user tier.
                - Create transaction including buying info.
                - Call payment to debit user balance (coming soon).
        - Allocate Voucher (Internal call):
            - (1) Check if exists valid voucher based on campaign id. Valid voucher must be available status and not ended yet.
            - (2) Create voucher for user in transactional operation:
                - Get voucher code from pool.
                - Create user voucher with Allocated status.
                - Update voucher configuration: increase allocated_qty, if allocated_qty >= max_qty => mark it Unavailable.
                - Delete used voucher code in pool.
        - Monitor voucher configuration (background task):
            - Run per 1 minutes by cronjob.
            - Mark Unavailable for voucher configuration has end_at < now.
        - Monitor user voucher (background task):
            - Run per 1 minutes by cronjob.
            - Mark Expired for voucher configuration has expire_at < now.
        - Provision voucher code pool (background task):
            - Run per 5 minutes by cronjob.
            - Generate voucher codes such that number of pool qty is equal to target.
                - (1) Generate missing qty by rules.
                - (2) Check if exists in the pool.
                - (3) Check if it's used by user.
                - (4) Continue go to (1) if not enough.
                - (5) Create generated codes for pool.