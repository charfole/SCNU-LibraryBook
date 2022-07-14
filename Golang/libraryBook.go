// @libraryBook.go
// @该程序用于预约华南师范大学图书馆的研讨室（六人间），请在校园网的环境下使用，并请勿将该程序用作恶意攻击网站、恶意霸占公共资源的手段！
// @charfole  2022-06-28

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/eddieivan01/nic"
)

// 下面是请求用到的常量
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"
var url = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/login.aspx"

// 获取后天的日期，如果想预约明天的研讨室，下面要改成AddDate(0, 0 ,1)
var date = time.Now().AddDate(0, 0, 2).String()
var houtian = strings.Fields(date)[0]

// 分早中晚时段，分别存放预约用到的链接
var MorningSlice, AfternoonSlice, NightSlice []string

// 传入预约用户的accno参数，以及早中晚的预约起始时间与结束时间，从而获取预约链接
func getBookLinks(user1, user2, user3, morningStart, morningEnd, afternoonStart, afternoonEnd, nightStart, nightEnd string) ([]string, []string, []string) {
	var Morning1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
	var Morning2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
	var Morning3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
	var Morning4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"

	var Afternoon1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
	var Afternoon2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
	var Afternoon3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
	var Afternoon4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"

	var Night1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
	var Night2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
	var Night3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
	var Night4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"

	s1 := []string{Morning1, Morning2, Morning3, Morning4}
	s2 := []string{Afternoon1, Afternoon2, Afternoon3, Afternoon4}
	s3 := []string{Night1, Night2, Night3, Night4}
	return s1, s2, s3
}

// 登录账户，获取用户的accno参数以及session，用于后续的请求
func login(id string, pwd string) (string, *nic.Session) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"
	url := "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/login.aspx"
	session := nic.NewSession()
	resp, _ := session.Post(url, &nic.H{
		Data: nic.KV{
			"id":  id,
			"pwd": pwd,
			"act": "login",
		},
		Headers: nic.KV{
			"Referer":    "http://lib-ic.scnu.edu.cn/ClientWeb/xcus/ic2/Default.aspx",
			"User-Agent": userAgent,
		},
	})
	fmt.Println(resp.Text)
	fmt.Println()
	// 分割返回的body，获取用户的accno
	s := strings.Split(resp.Text, ",")
	// 冒号分割键值对，获取值
	accno := strings.Split(s[4], ":")[1]
	// 去除双引号
	accno = accno[1 : len(accno)-1]
	return accno, session
}

// 对早上的研讨室进行预约
func morning(wg *sync.WaitGroup, session *nic.Session, testFlag string) {
	// 当函数完成后，调用wg.Done()进行协程同步
	defer wg.Done()
	MorningSession := session
	var startIndex int
	if testFlag == "优先约早上" {
		startIndex = 1
	} else {
		startIndex = 0
	}
	for i := startIndex; i < len(MorningSlice); i++ {
		resp, _ := MorningSession.Get(MorningSlice[i], nil)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			fmt.Println("Morning:", i)
			fmt.Println(resp.Text)
			return
		} else {
			fmt.Println("Morning:", i)
			fmt.Println(resp.Text)
			continue
		}
	}
	return
}

// 对下午的研讨室进行预约
func afternoon(wg *sync.WaitGroup, session *nic.Session, testFlag string) {
	// 当函数完成后，调用wg.Done()进行协程同步
	defer wg.Done()
	AfternoonSession := session
	var startIndex int
	if testFlag == "优先约下午" {
		startIndex = 1
	} else {
		startIndex = 0
	}
	for i := startIndex; i < len(AfternoonSlice); i++ {
		resp, _ := AfternoonSession.Get(AfternoonSlice[i], nil)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			fmt.Println("Afternoon:", i)
			fmt.Println(resp.Text)
			return
		} else {
			fmt.Println("Afternoon:", i)
			fmt.Println(resp.Text)
			continue
		}
	}
	return
}

// 预约晚上的研讨室
func night(wg *sync.WaitGroup, session *nic.Session, testFlag string) {
	// 当函数完成后，调用wg.Done()进行协程同步
	defer wg.Done()
	NightSession := session
	var startIndex int
	if testFlag == "优先约晚上" {
		startIndex = 1
	} else {
		startIndex = 0
	}
	for i := startIndex; i < len(NightSlice); i++ {
		resp, _ := NightSession.Get(NightSlice[i], nil)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			fmt.Println("Night:", i)
			fmt.Println(resp.Text)
			return
		} else {
			fmt.Println("Night:", i)
			fmt.Println(resp.Text)
			continue
		}
	}
	return
}

// 嗅探函数，判断是否可以开始预约
func sniff(Session *nic.Session, morningBook int, afternoonBook int, nightBook int) (string, bool) {
	var resp *nic.Response
	resp = new(nic.Response)
	var testFlag = "0"
	var bookFlag = false

	// 用Session的人做嗅探
	if nightBook == 1 {
		testFlag = "优先约晚上"
		resp, _ = Session.Get(NightSlice[0], nil)

		fmt.Println("Night:", 0)
		fmt.Println("嗅探结果：", resp.Text)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			bookFlag = true
		}
	} else if afternoonBook == 1 && testFlag == "0" {
		testFlag = "优先约下午"
		resp, _ = Session.Get(AfternoonSlice[0], nil)
		fmt.Println("Afternoon:", 0)
		fmt.Println("嗅探结果：", resp.Text)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			bookFlag = true
		}
	} else if morningBook == 1 && testFlag == "0" {
		testFlag = "优先约早上"
		resp, _ = Session.Get(MorningSlice[0], nil)
		fmt.Println("Morning:", 0)
		fmt.Println("嗅探结果：", resp.Text)

		respSplit := strings.Split(resp.Text, ",")
		msg := strings.Split(respSplit[2], ":")[1]

		if msg == "\"操作成功！\"" {
			bookFlag = true
		}
	}

	// 逗号分割出返回body中的 "msg":"xxxxx" 字段
	respSplit := strings.Split(resp.Text, ",")
	//当返回的信息有两个冒号，比如出现20:00，需要特殊处理，此时msg的长度为3
	msg := strings.Split(respSplit[2], ":")
	// 还没到八点的情况
	if len(msg) == 3 {
		return "还没到八点", bookFlag
	} else {
		return testFlag, bookFlag
	}
}

// 预约函数，用于对早中晚的研讨室分别发起预约
func book(SessionOne *nic.Session, SessionTwo *nic.Session, SessionThree *nic.Session, testFlag string, morningBook int, afternoonBook int, nightBook int) {
	var bookAll = morningBook + afternoonBook + nightBook
	var wg sync.WaitGroup
	wg.Add(bookAll)

	if morningBook == 1 {
		go morning(&wg, SessionOne, testFlag)
	}
	if afternoonBook == 1 {
		go afternoon(&wg, SessionTwo, testFlag)
	}
	if nightBook == 1 {
		go night(&wg, SessionOne, testFlag)
	}
	// 等待所有任务完成
	wg.Wait()
}

func main() {
	// 传入学号和图书馆账号的密码，获取账户的accno和session
	var usr1, SessionOne = login("账户1", "密码1")
	var usr2, SessionTwo = login("账户2", "密码2")
	var usr3, SessionThree = login("账户3", "密码3")

	// 根据用户的accno，与早中晚研讨室预约所对应的时间段（早上08：30-11：30，下午14：00-17：30，晚上19：00-21：30），来获取对应的研讨室预约链接
	MorningSlice, AfternoonSlice, NightSlice = getBookLinks(usr1, usr2, usr3, "0830", "1130", "1400", "1730", "1900", "2130")

	var onTime string
	var bookFlag bool

	// 下面三个值分别对应早中晚研讨室的预约状态，1为需要预约，0为不需要
	var morningBook int = 1
	var afternoonBook int = 1
	var nightBook int = 1

	for {
		// 为了防止同一个人，在间隔不到4小时内连续预约两次研讨室而导致冲突，需要进行特殊处理
		if nightBook == 1 {
			onTime, bookFlag = sniff(SessionOne, morningBook, afternoonBook, nightBook)
		} else {
			onTime, bookFlag = sniff(SessionTwo, morningBook, afternoonBook, nightBook)
		}
		fmt.Println(onTime)
		fmt.Println()
		if onTime == "还没到八点" {
			continue
		} else {
			//最后三个参数为是否预约早中晚
			fmt.Println("嗅探完毕，开始预约")
			start := time.Now()
			if bookFlag == true {
				if onTime == "优先约晚上" {
					// sniff函数中成功预约晚上，不需要在主函数中预约晚上
					book(SessionOne, SessionTwo, SessionThree, onTime, 1, 1, 0)
				}
				if onTime == "优先约下午" {
					// sniff函数中成功预约下午，不需要在主函数中预约下午
					book(SessionOne, SessionTwo, SessionThree, onTime, 1, 0, 0)
				}
				if onTime == "优先约早上" {
					// sniff函数中成功预约早上，不需要在主函数中预约早上
					book(SessionOne, SessionTwo, SessionThree, onTime, 0, 0, 0)
					break
				}
			} else if bookFlag == false { //嗅探函数预约不成功，进行正常预约
				book(SessionOne, SessionTwo, SessionThree, onTime, morningBook, afternoonBook, nightBook)
			}
			end := time.Since(start)
			fmt.Println("Time cost: ", end)
			fmt.Println()
		}
	}
}
