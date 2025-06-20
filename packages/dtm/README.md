```sql
create
database dtm_busi default character set utf8mb4 collate utf8mb4_unicode_ci;
use dtm_busi;
CREATE TABLE dtm_busi.`user_account`
(
    `id`              int(11) AUTO_INCREMENT PRIMARY KEY,
    `user_id`         int(11) not NULL UNIQUE,
    `balance`         decimal(10, 2) NOT NULL DEFAULT '0.00',
    `trading_balance` decimal(10, 2) NOT NULL DEFAULT '0.00',
    `create_time`     datetime                DEFAULT now(),
    `update_time`     datetime                DEFAULT now()
);

```