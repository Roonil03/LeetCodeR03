class Solution {
    public long dividePlayers(int[] skill) {
        // Map<Integer, Integer> hmap = new HashMap<>();
        // for(int i = 0; i<skills.length;i++)
        // {
        //     if(hmap.containsKey(target - skills[i])){
        //         return new int[] {hmap.get(target-nums[i]),i};
        //     }
        //     hmap.put(nums[i],i);
        // }
        Arrays.sort(skill);
        int sum = 0;
        int pair[][] = new int[2][skill.length/2];
        pair[0][0] = skill[0];
        pair[1][0] = skill[skill.length-1];
        sum = pair[0][0] + pair[1][0];
        for(int i = 1;i<skill.length/2;i++)
        {
            pair[0][i] = skill[i];
            pair[1][i] = skill[skill.length-i-1];
            if(sum != (pair[0][i] + pair[1][i]))
            {
                return -1;
            }
        }
        long chemistry = 0;
        for(int i = 0; i<pair[1].length;i++)
        {
            chemistry += (pair[0][i] * pair[1][i]);
        }
        return chemistry;
    }
}