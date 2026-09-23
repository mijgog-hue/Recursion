CodeWars:Well of Ideas - Easy Version
Time: O(n)
Space: O(1)

package kata

func Well(x []string) string {
  good := 0
  for _, v:= range x {
    if v == "good"{
     good++
      }
    }
   if good ==0 {
         return "Fail!"
      } else if good <= 2{
 return "Publish!"  
   }
  return "I smell a series!"
  }
  
