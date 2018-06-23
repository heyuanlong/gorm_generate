# gorm_generate

gorm_generate.exe -s beegoblog -t users

github.com/jinzhu/gorm

```
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `articleid` int(11) NOT NULL,
  `content` text,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名字',
  `time` varchar(19) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '时间',
  PRIMARY KEY (`id`),
  KEY `articleid` (`articleid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='评论';


type Comment struct {
        ID     uint64  `gorm:"primary_key" json:"-"`
        Articleid int`gorm:"column:articleid" json:"articleid"`
        Content string`gorm:"column:content" json:"content"`
        Name string`gorm:"column:name" json:"name"`
        Time string`gorm:"column:time" json:"time"`
}
func (Comment) TableName() string {
        return "comment"
}
```