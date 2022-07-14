# SCNU-LibraryBook
该脚本用于预约华南师范大学图书馆研讨室的脚本，仅需输入预约用户的账户和密码，运行脚本，便可对研讨室发起预约。目前脚本仅支持预约六人的研讨室，如需预约其他类型的研讨室，需要获取对应的预约链接并修改预约成员。

## 使用方法

使用脚本时必须处在**校园网**环境下，否则会被认证页面拦截。图书馆的预约规则为第 x 天晚上八点开始预约第 x + 2 天的研讨室，因此预约的日期默认为后天，如需修改为当天或者明天，可以看下面的说明。Golang 和 Python 脚本的预约逻辑和效果大致一样，都是生成一条符合规则的预约链接 post 到预约的 url，不同点在于 Golang 的预约脚本中使用了协程，可以并行地发起多个预约，因此预约速度会更快。

**1.clone 本仓库**

```
git clone https://github.com/charfole/SCNU-LibraryBook.git
```

**2.修改预约信息**

**Golang 版本**

```
# 修改日期，预约明天的研讨室，下面要改成AddDate(0, 0 ,1)，当天则改成AddDate(0, 0 ,0)
var date = time.Now().AddDate(0, 0, 2).String()

# 主函数中，填入三个预约人的图书馆研讨室系统的账户和密码
var usr1, SessionOne = login("账户1", "密码1")
var usr2, SessionTwo = login("账户2", "密码2")
var usr3, SessionThree = login("账户3", "密码3")

# 主函数中，填入预约早中晚研讨室的时间（以下的链接对应为：早上08：30-11：30，下午14：00-17：30，晚上19：00-21：30）
MorningSlice, AfternoonSlice, NightSlice = getBookLinks(usr1, usr2, usr3, "0830", "1130", "1400", "1730", "1900", "2130")

# 主函数中，下面三个值分别对应早中晚研讨室的预约状态，1为需要预约，0为不需要
var morningBook int = 1
var afternoonBook int = 1
var nightBook int = 1
```

**Python 版本**

```
# 修改日期，预约明天的研讨室，下面要改成datetime.timedelta(days=1)，当天则改成datetime.timedelta(days=0)
houtian = str(today + datetime.timedelta(days=2))

# 主函数中，填入三个预约人的图书馆研讨室系统的账户和密码
user1, FirstBookSession = LibraryLogin("账户1", "密码1")
user2, SecondBookSession = LibraryLogin("账户2", "密码2")
user3, ThirdBookSession = LibraryLogin("账户3", "密码3")

# 主函数中，填入预约早中晚研讨室的时间（以下的链接对应为：早上08：30-11：30，下午14：00-17：30，晚上19：00-21：30）
urlMorning, urlAfternoon, urlNight = getBookLinks(user1, user2, user3, "0830", "1130", "1400", "1700", "1900", "2130")

# 主函数中，下面三个值分别对应早中晚研讨室的预约状态，1为需要预约，0为不需要
morningBook, afternoonBook, nightBook = 1, 1, 1
```

**2.运行 Golang 脚本**

```
cd Golang
go run libraryBook.go
# 预约成功后，Ctrl+c 停止脚本运行即可
```

**3.运行 Python 脚本**

```
cd Python
python3 libraryBook.py
# 预约成功后，Ctrl+c 停止脚本运行即可
```

