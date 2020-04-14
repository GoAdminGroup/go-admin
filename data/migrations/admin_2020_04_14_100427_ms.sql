CREATE TABLE[goadmin_site] (
 [id] int   identity(1,1) ,
 [key] varchar(100)   NOT NULL,
 [value] text   NOT NULL,
 [state] tinyint   NOT NULL DEFAULT 0,
 [description] varchar(3000)   NOT NULL,
 [created_at] datetime NULL DEFAULT GETDATE(),
 [updated_at] datetime NULL DEFAULT GETDATE(),
  PRIMARY KEY ([id]),
)
