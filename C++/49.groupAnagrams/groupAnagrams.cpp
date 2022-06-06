/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

 

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
#include<vector>
#include<algorithm>
#include<unordered_map>
using namespace std;
class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string,int> records;
        vector<vector<string>> res;
        int n = strs.size();
        int pos = 0;
        for(int i = 0; i < n; i++){
            string s = strs[i];
            sort(s.begin(),s.end());
            if(records.find(s) != records.end()){
                //如果strs[i]在中
                res[records[s]].push_back(strs[i]);
            }else{
                records[s] = pos++;
                vector<string> string_set;
                res.push_back(string_set);
                res[records[s]].push_back(strs[i]);
            }
        }
        return res;
    }
};