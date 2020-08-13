package timedifferencedraft

import (
	//"fmt"
	"strconv"
	"time"
)
func TimeDifferencedraft(postdate string)string{
	loc,_:=time.LoadLocation("UTC")
	dt:=time.Now().In(loc)
	ct:=dt.Format("2006-01-02 3:04:05")
	x,_:=time.Parse("2006-01-02 3:04:05",postdate)
	y,_:=time.Parse("2006-01-02 3:04:05",ct)
	diff:=y.Sub(x)
	//fmt.Println(int(diff.Hours()))
	timediff:=int(diff.Hours())
	if timediff<=24{
		if timediff<1{
			min:=diff.Hours()*60
			if int(min)==1{
			minutes:=strconv.Itoa(int(min))
			return "Created "+minutes+" minute ago"
			}
			if int(min)>1{
				minutes:=strconv.Itoa(int(min))
				return "Created "+minutes+" minutes ago"
				}
			if int(min)<1{
				return "Created a few seconds ago"
			}
		}
		if timediff==1{
			hours:=strconv.Itoa(timediff)
			return "Created "+hours+" hour ago"
		}
		if timediff>1{
			hours:=strconv.Itoa(timediff)
			return "Created "+hours+" hours ago"
		}
	}
	if timediff>24{
		days:=timediff/24
		if days==1{
			nofdays:=strconv.Itoa(days)
			return "Created "+nofdays+" day ago"
		}
		if days<=30&&days!=1{
		nofdays:=strconv.Itoa(days)
		return "Created "+nofdays+" days ago"
	}
	if days>30{
		months:=days/30
		if months==1{
			nofmonths:=strconv.Itoa(months)
			return "Created "+nofmonths+" month ago"
			}
		if months<12&&months!=1{
		nofmonths:=strconv.Itoa(months)
		return "Created "+nofmonths+" months ago"
		}
		if months>=12{
			years:=months/12
			nofyears:=strconv.Itoa(years)
			if years==1{
				
			return "Created "+nofyears+" year ago"
			}
			if years>1{
				return "Created "+nofyears+" years ago"
			}
		}
	}
	}
	return "error"
	}
