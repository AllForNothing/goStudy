package split_string

import "strings"

func Split(str, sep string)(ret []string)  {
   index := strings.Index(str, sep)
   if index >= 0 {
	   ret = append(ret, str[:index])
	   str = str[index+1:]
	   index = strings.Index(str, sep)
   }
   ret = append(ret, str)
   return
}
