package first;

import java.util.ArrayList;

public class Solution {
    /**
     * 失物排序方法
     * @param lostArray 待排序的失物数组
     */
    public void sortLost(LostThing[] lostArray){
        for (int i = 0; i < lostArray.length-1; i++) {
            for (int j = i+1; j < lostArray.length; j++) {
                if (lostArray[i].getLostTime().compareTo(lostArray[j].getLostTime()) == -1){
                    LostThing tmp =lostArray[i];
                    lostArray[i]=lostArray[j];
                    lostArray[j]=tmp;
                }
            }
        }
    }

    /**
     * 按关键字搜索失物的方法，这里假设按照失物的领取地点进行搜索
     * @param lostArray 失物数组
     * @param keyword 用户输入的关键字
     * @return 返回查找到的失物
     */
    public LostThing[] selectByKeyword(LostThing[] lostArray,String keyword){
        ArrayList<LostThing> list =new ArrayList<>();
        for (LostThing lost: lostArray) {
            if (lost.getToAddress().contains(keyword)){
                list.add(lost);
            }
        }
        if (list.isEmpty()){
            return null;
        }
        return list.toArray(new LostThing[0]);
    }

}
