#include<vector>
#include<algorithm>
using namespace std;
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution {
public:
    vector<int> temp;
    vector<vector<int>> permute(vector<int>& nums) {
        //使用visited数组方法
        int n = nums.size();
        sort(nums.begin(), nums.end());
        vector<int> visited(n,0);
        vector<vector<int>> res;
        permute_(res,nums,visited);
        return res;
    }
    void permute_(vector<vector<int>>& res, vector<int>& nums, vector<int>& visited)
    {   
        if (temp.size() > nums.size())
        {
            return;
        }
        
        if (temp.size() == nums.size())
        {
            res.emplace_back(temp);
            temp.clear();
        }
        
        for (int i = 0; i < nums.size(); i++)
        {
            if (visited[i])
            {
                continue;
            }
            
            visited[i] = 1;
            temp.emplace_back(nums[i]);
            permute_(res,nums,visited);
            visited[i] = 0;

        }
        return;
    }
};