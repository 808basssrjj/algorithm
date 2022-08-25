## 1. 引入

Manacher算法和KMP算法都是解决字符串相关题目的常见算法原型，但是各自解决的问题却不一样。

Manacher算法一开始是专门用来解决 "字符串中最长回文子串问题" 的，实际上Manacher算法中有一个非常重要的信息，使用它可以解决很多其他问题，这个信息就是：**每一个位置的最长回文半径**



## 2. 经典解法

将字符串中的每一位字符都当作对称轴，然后同时向左右两边开始匹配。这种方法有一个明显的问题，就是如果回文子串的长度是偶数，那么是没有办法检测出来的，因为长度是偶数的回文串实际上对称轴是 "虚" 的，无法实际定位。

假设当前字符串是 "122131221"。

如果使用我们原来脑补的方法，则 "1221" 回文子串就不会被检测到。

我们将原来脑补的方法进行改进，在匹配前先对原始字符串进行处理。处理方式是：字符串左右两头加 "#"，每两个字符中间加 "#"。处理成："#1#2#2#1#3#1#2#2#1#"。然后将处理之后的字符串中的每一个字符都当作对称轴，同时向左右两边开始匹配，记录回文子串长度。最后将每个位置统计的回文子串长度除以2（向下取整），对应到原字符串就是以该位置为对称轴的回文子串的长度。

```go
func manacherString(str string) []rune {
	chars := []rune(str)
	res := make([]rune, len(chars)*2+1)
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			res[i] = '#'
		} else {
			res[i] = chars[i>>1]
		}
	}
	return res
}
```

使用改进后的方法，无论是奇数个数的回文子串还是偶数个数的回文子串都可以被检测到。

我们想一个问题：**处理原字符串时加入的辅助字符要不要求是原字符串中没有出现的字符？**

辅助字符是什么都行，不会影响最后的答案。因为无论是以处理后的字符串的哪一个字符作为对称轴向左向右开始匹配，永远都是辅助字符和辅助字符比，真实字符和真实字符比。



## 3. 回文半径和回文直径

在讲Manacher算法之前，我们需要了解回文半径和回文直径的概念。

**回文直径**：从对称轴开始向左向右衍生，直到回文区域边界后统计的字符总数。

**回文半径**：从对称轴开始向左或向右衍生，直到回文区域边界后统计的字符总数。



## 4. Manacher算法中的基础概念

Manacher和经典解法的处理流程实际上是一样的，Manacher和KMP一样主要是设计了加速操作，Manacher能够将时间复杂度优化到O(N)。

Manacher算法设计中有三个重要点：

### 1.  回文半径数组radius

​	回文半径数组radius是用来记录以每个位置的字符为回文中心求出的回文半径长度

​	#a#c#b#b#c#b#d#s#的回文半径数组为[1, 2, 1, 2, 1, 2, 5, 2, 1, 4, 1, 2, 1, 2, 1, 2, 1]

### 2. 最右回文右边界R

​	记录之前匹配回文区域的所有字符中，回文边界达到的最右下标（初值为-1）

​	#a#c#b#b#c#b#d#s#

​	最开始的时候R=-1，到p=0的位置，回文就是其本身，最右回文右边界R=0;p=1时，有回文串#a#，R=2；p=2时，R=2;P=3时，R=6;p=4时，最右回	文右边界还是p=3时的右边界，R=6,依次类推

​	R 永远是只增不减的，只要有字符的回文更靠右，下标就会更新给 R。

### 3. 最右回文右边界的对称中心C

​	和 R 一起用，记录当前取得最右下标的回文区域的中心点的下标（初值为-1，如果最右下标重合，按照原中心点的下标）

​	C 也是永远只增不减的，R 更新 C 一定更新，R 不更新 C 一定不更新。



## 5. 加速流程

首先需要构建回文半径数组，在构建回文半径数组时，会遇到两种大情况：

**第一种大情况**：当前匹配的字符的位置不在之前匹配的字符的回文区域的最右边界中。该情况无优化，只能从该中心点开始同时向两边暴力扩展匹配，同时计算出该字符的回文半径。

例如：

![20211014100421.png](C:\Users\Administrator\Desktop\study\algorithm\note\img\manacher.png)



**第二种大情况**：当前匹配的字符的位置在之前匹配的字符的回文区域的最右边界中（如上图当 R=2 时的情况）。

当第二种大情况出现的时候，一定存在下图表示的通用拓扑结构：

i 为当前匹配的字符的位置，i‘ 是以 C 为对称轴所作的 i 的对称点。C、L 和R 一定都存在。

![](C:\Users\Administrator\Desktop\study\algorithm\note\img\manacher2.png)

> i 和 C 是不可能重合的，因为 C 表示的是 i 之前字符构建的最长回文子串的中心点。当遍历到 i 位置时，C 一定已经遍历过了。

按照 i’ 的回文区域的状况可以将第二种大情况划分成三种具体的情况，每一种情况都有单独的优化。



**（1）第一种情况**：i‘ 的回文区域完全在L~R的内部

![](C:\Users\Administrator\Desktop\study\algorithm\note\img\mn1.png)

此时，i 的回文半径就是 i’ 的回文半径。

证明:  设置 a~b 区域前一个字符为 X，后一个字符为 Y；设置 c~d 区域前一个字符为 Z，后一个字符为 K。

- 因为 i 和 i' 关于C对称  所以 i 的回文范围至少是(Z,K)   且 Y==Z  X==K
- 因为 i' 的回文范围为[a,b],所以 X!=Y , 所以 Z!=L



**（2）第二种情况**：i‘ 的回文区域有一部分在 L~R 的外部。即 iL(i'的左边界) < L

![](C:\Users\Administrator\Desktop\study\algorithm\note\img\mn2.png)

此时，i 的回文半径就是 i~R 的距离。

证明: L 作 i’ 的对称点 L'，R 作 i 的对称点 R'。

- 因为 X 和 Y 都在 i‘ 的回文区域中，且关于 i’ 对称，所以X == Y。

- 又因为 Y 和 Z 都在 C 的回文区域中，且关于 C 对称，所以Y == Z。
- 又因为 C 的回文为 L 到 R , 所以 X != K
- 所以 Z != K



**（3）第三种情况**：i‘ 的回文区域的左边界正好和 L 重合。

![20211014173905.png](C:\Users\Administrator\Desktop\study\algorithm\note\img\mn3.png)

此时，不能直接得出 i 的回文半径，**只能确定回文半径最小就是 i~R**。能不能更大，需要从 i 位置向外扩展匹配。



## 6. 代码实现

```go
// Manacher1 分支判断版
func Manacher1(str string) int {
	if len(str) < 0 {
		return 0
	}
	chars := manacherString(str)

	res := -1 << 31
	radius := make([]int, len(chars)) // 回文半径数组
	R, C := -1, -1                    // 最右边界和中心
	for i := 0; i < len(chars); i++ {
		if i > R { // 一.i在R外部:从i开始暴力扩, R,C更新
			//半径至少为1
			radius[i] = 1
			left, right := i-1, i+1
			for left > -1 && right < len(chars) && chars[left] == chars[right] {
				radius[i]++ //半径+1
				left--
				right++
			}
			//R,C更新
			if right-1 > R {
				R = right - 1
				C = i
			}
		} else { // 二.i在R内部
			i2 := 2*C - i //i2: i关于C对称的位置
			// 左边界 = mid-radius+1
			iL := i2 - radius[i2] + 1 //i'的左边界
			L := C - radius[C] + 1    //C的左边界
			if iL > L {               // 2.1 i'回文范围在L..R内 : i的答案就是i'的答案
				//radius[i] = radius[2*C-i]
				radius[i] = radius[i2]
			} else if iL < L { // 2.2 i'回文范围有一部分在L..R外  : i的答案就是i到R的距离
				radius[i] = R - i + 1
			} else { // 2.3 i'的回文左边界压线
				//范围至少为[R',R],半径为i到R的距离
				radius[i] = R - i + 1
				left, right := i-radius[i], R+1
				//left, right := 2*i-R-1, R+1 // 也可以按中心对称求
				for ; left > -1 && right < len(chars) && chars[left] == chars[right]; left, right = left-1, right+1 {
					radius[i]++ //半径+1
				}
				//R,C更新
				if right-1 > R {
					R = right - 1
					C = i
				}
			}
		}
		if radius[i] > res {
			res = radius[i]
		}
	}
	fmt.Println(radius)
	return res - 1 //回文长度=回文半径-1(因为一开始加了字符)
}

// Manacher 改进版
func Manacher(str string) int {
	if len(str) < 0 {
		return 0
	}
	chars := manacherString(str)

	res := -1 << 31
	radius := make([]int, len(chars)) // 回文半径数组
	R, C := -1, -1                    // R:最右边界再往右一个位置,最右有效范围是R-1位置  C:R边界的中心位置
	for i := 0; i < len(chars); i++ {
		// i至少的回文区域,先给radius
		if i >= R {
			radius[i] = 1
		} else {
			radius[i] = min(radius[2*C-i], R-i)
		}
		// 然后向两边扩
		left, right := i-radius[i], i+radius[i]
		for left > -1 && right < len(chars) {
			if chars[left] == chars[right] {
				radius[i]++ //半径+1
				left, right = left-1, right+1
			} else {
				break
			}
		}
		// 更新C,R
		if i+radius[i] > R {
			R = i + radius[i]
			C = i
		}

		res = max(radius[i], res)
	}
	fmt.Println(radius)
	return res - 1 //回文长度=回文半径-1(因为一开始加了字符)
}
```



## 7. 复杂度

```java
public static int process(char[] str) {
    int R = ?;
    int C = ?;

    int[] next = new int[str.length];

    for (int i = 0; i < str.length; i ++) {
        if (i在R的外面) {
            从i开始向外扩;
        } else { // i在R的里面
            if (i'的回文区域完全在L~R中) {
                // O(1)的操作
                返回i'的回文半径; 
            } else if (i'的回文区域的左边界在L的左边) {
                // O(1)的操作
                返回R-i; 
            } else { // i'的回文区域的左边界与L重合
                从R开始向外扩;
        }
    }
 	                      
    排序获得最长的回文半径处理成原文的回文串长度返回 
}
```

第一种大情况，和第二种大情况的第③种小情况需要向外扩充匹配，因此必定会失败1次。

第二种大情况的第①种小情况和第②种小情况是不需要向外扩充匹配的，因此失败0次。

因此每个位置扩充失败的代价是O(N)。

根据上述伪代码，不看失败，只看成功。假设字符串长度为N，以 i 和 R 作为参考标注：

|            | i1 (max -> N) | i1-i2 (max -> N) |
| ---------- | ------------- | ---------------- |
| 第一个分支 | 增大          | 增大             |
| 第二个分支 | 增大          | 不变             |
| 第三个分支 | 增大          | 不变             |
| 第四个分支 | 增大          | 增大             |

i 只增不减，R 也是只增不减。

我们将扩的行为和 R 变大绑定到一起，每一次扩，R 都会变大，R变大的总幅度就是扩成功的次数。而扩失败的次数，是可以估计出来一个总的量，就是O(N)。

因此整个时间复杂度为O(N)。

