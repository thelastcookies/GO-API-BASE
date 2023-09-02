# model

`Model` 层，或者叫 `Entity`、实体层。

用于存放实体类，为数据库表的映射。

相关约定：

- 默认使用 `MySQL` 数据库。

- 使用 `GORM` 作为 ORM 库。

- 声明实体类的结构体符合 `GORM` 的约定规范。

- 一个表中需要包含的三大字段：
  - 主键(id)
  - 创建时间(created_at)
  - 更新时间(updated_at)