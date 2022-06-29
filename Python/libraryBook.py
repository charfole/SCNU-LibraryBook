# @libraryBook.py
# @该程序用于预约华南师范大学图书馆的研讨室（六人间），请在校园网的环境下使用，并请勿将该程序用作恶意攻击网站、恶意霸占公共资源的手段！
# @charfole  2022-06-29

# -*- coding: utf-8 -*-

import requests
import time
import datetime
import json

# 获取后天的日期，如果想预约明天的研讨室，下面要改成datetime.timedelta(days=1))
today = datetime.date.today()
houtian = str(today + datetime.timedelta(days=0))

# 下面是请求用到的常量
userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"
Header = {
    "Referer": "http://lib-ic.scnu.edu.cn/ClientWeb/xcus/ic2/Default.aspx",
    'User-Agent': userAgent,
}

# 传入预约用户的accno参数，以及早中晚的预约起始时间与结束时间，从而获取预约链接
def getBookLinks(user1, user2, user3, morningStart, morningEnd, afternoonStart, afternoonEnd, nightStart, nightEnd): 
    Morning_1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
    Morning_2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
    Morning_3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
    Morning_4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + morningStart[0:2] + "%3A" + morningStart[2:] + "&end=" + houtian + "+" + morningEnd[0:2] + "%3A" + morningEnd[2:] + "&start_time=" + morningStart + "&end_time=" + morningEnd + "&act=set_resv&_=1605268746419"
    
    Afternoon_1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
    Afternoon_2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
    Afternoon_3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"
    Afternoon_4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + afternoonStart[0:2] + "%3A" + afternoonStart[2:] + "&end=" + houtian + "+" + afternoonEnd[0:2] + "%3A" + afternoonEnd[2:] + "&start_time=" + afternoonStart + "&end_time=" + afternoonEnd + "&act=set_resv&_=1605268746419"

    Night_1 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455538&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
    Night_2 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455542&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
    Night_3 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455550&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
    Night_4 = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/reserve.aspx?dialogid=&dev_id=100455554&lab_id=100455305&kind_id=100455303&type=dev&min_user=3&max_user=6&mb_list=%24" + user1 + "%2C" + user2 + "%2C" + user3 + "&start=" + houtian + "+" + nightStart[0:2] + "%3A" + nightStart[2:] + "&end=" + houtian + "+" + nightEnd[0:2] + "%3A" + nightEnd[2:] + "&start_time=" + nightStart + "&end_time=" + nightEnd + "&act=set_resv&_=1605268746419"
     
    urlMorning = [Morning_4,Morning_3,Morning_2,Morning_1]
    urlAfternoon = [Afternoon_4,Afternoon_3,Afternoon_2,Afternoon_1]
    urlNight = [Night_4,Night_3,Night_2,Night_1]
    return urlMorning, urlAfternoon, urlNight

# 登录账户，获取用户的accno参数以及session，用于后续的请求
def LibraryLogin(account, password):
    session = requests.Session()
    postUrl = "http://lib-ic.scnu.edu.cn/ClientWeb/pro/ajax/login.aspx"
    postData = {
        "id": account,
        "pwd": password,
        "act": "login",
    }
    responsePost = session.post(postUrl, data = postData, headers = Header)
    # 无论是否登录成功，状态码一般都是 statusCode = 200
    accno = json.loads(responsePost.text)["data"]['accno']
    print(f"statusCode = {responsePost.status_code}")
    print(f"text = {responsePost.text}")
    print(f"accno: {accno}")
    return accno, session

# 预约函数，用于对早中晚的研讨室分别发起预约
def LibraryBook(session_1, session_2, session_3,urlMorning, urlAfternoon, urlNight, morning, afternoon, night):
    start = time.time()
    test_flag = '0'

    if(night==1):
        response = session_3.get(urlNight[0],headers=Header).json()
        test_flag = '优先约晚上'
    if(afternoon==1 and test_flag=='0'):
        response = session_3.get(urlAfternoon[0],headers=Header).json()
        test_flag = '优先约下午'
    if(morning==1 and test_flag=='0'):
        response = session_3.get(urlMorning[0],headers=Header).json()
        test_flag = '优先约早上'
    
    msg = response['msg']
    print(msg)

    # 先预约一个研讨室，判断是否到达可预约时间
    if (msg == "操作成功"):
        if(test_flag == '优先约晚上'):
            night = 0
        elif(test_flag == '优先约下午'):
            afternoon = 0
        elif(test_flag == '优先约早上'):
            morning = 0
    if (msg[-4:] == '方可预约'):
        print('还没到8点')
        return

    # 对早中晚的研讨室进行预约
    else:
        print()
        print("------------------------")

        if(night == 1):
            print("开始约晚上")
            numOfNight = len(urlNight) - 1
            for i in range(numOfNight): # 这里只有numOfNight减一是因为上面在嗅探成功时，已尝试过预约某个研讨室，因此后续只需要预约后面的几个研讨室
                response = session_1.get(urlNight[i+1],headers=Header).json()
                msg_1 = response['msg']
                print(msg_1)
                if (msg_1 == "操作成功"):
                    break
            print()
        
        if(afternoon == 1):
            print("开始约下午")
            if(test_flag == '优先约下午'):
                numOfAfternoon = len(urlAfternoon) - 1
            else:
                numOfAfternoon = len(urlAfternoon)
            for i in range(numOfAfternoon):
                response = session_2.get(urlAfternoon[i],headers=Header).json()
                msg_1 = response['msg']
                print(msg_1)
                if (msg_1 == "操作成功"):
                    break
            print()

        if(morning == 1):
            print("开始约早上")
            if(test_flag == '优先约早上'):
                numOfMorning = len(urlMorning) - 1
            else:
                numOfMorning = len(urlMorning)
            for i in range(numOfMorning):
                response = session_1.get(urlMorning[i],headers=Header).json()
                msg_1 = response['msg']
                print(msg_1)
                if (msg_1 == "操作成功"):
                    break
            print()

        print("------------------------")
        end = time.time()
        print (end-start)

if __name__ == "__main__":
    # 对需要用到的账户进行登录
    user1, FirstBookSession = LibraryLogin("账户1", "密码1")
    user2, SecondBookSession = LibraryLogin("账户2", "密码2")
    user3, ThirdBookSession = LibraryLogin("账户3", "密码3")
    urlMorning, urlAfternoon, urlNight = getBookLinks(user1, user2, user3, "0830", "1130", "1400", "1700", "1900", "2130")

    # 早中晚研讨室的预约情况, 0为不需要预约, 1为需要预约
    morningBook, afternoonBook, nightBook = 1, 1, 1

    # 不断进行预约
    while True:
        LibraryBook(FirstBookSession, SecondBookSession, ThirdBookSession, urlMorning, urlAfternoon, urlNight, morningBook, afternoonBook, nightBook)
        # 参数分别为三个用户的Session，研讨室的预约链接，以及是否约早中晚的mask, 1为需要预约, 0为不需要