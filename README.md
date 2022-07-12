# written-examination-LRU

实现一个缓存Cache， 缓存Cache会在其已满时驱逐最近最少使用的数据
例如：
缓存大小 ： Size 3
输入: Data ("key1", "1"), Data ("key2", "2"), Data ("key3", "3"),  Get "key1",  Put ("key4", "4")
输出: [("key4", "4"), ("key1", "1"), ("key3", "3")]
