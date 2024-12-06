# Problem assumption

Assume that we have the system managing users and provide subscription plans that users able to buy.
This source code help that system create campaign(s) and give first 100 users registered by campaign's link a voucher, this voucher will be discounted 30% when user subscribe SILVER plan.

---
# System architecture
![System architecture](./assets/system-architecture-overview.png "System architecture")

---
# Tech stack
- Considering by:
  - Scope of requirement.
  - Timeline: less than 3 days.
  - Scalability and Availability: able to run services on multiple pods.
  - Generality: able to support other campaigns.
- Details:
  - Microservices with Golang for highly performant services
  - Data management: MySQL for main database and Redis for data caching.

---