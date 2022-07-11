package first;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.math.BigInteger;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

class SolutionTest {
    Solution s=new Solution();
    @BeforeEach
    void setUp() {
        System.out.println("开始测试...");
    }

    @AfterEach
    void tearDown() {
        System.out.println("测试完毕");
    }

    @Test
    void sortLost() {

        LostThing thing1 = new LostThing(BigInteger.valueOf(1));
        LostThing thing2 = new LostThing(BigInteger.valueOf(2));
        LostThing thing3 = new LostThing(BigInteger.valueOf(3));
        LostThing thing4 = new LostThing(BigInteger.valueOf(4));
        LostThing[] expected = new LostThing[]{thing4,thing3,thing2,thing1};

        LostThing[] actual = new LostThing[]{thing4,thing1,thing3,thing2};
        s.sortLost(actual);

        for (int i = 0; i < 4; i++) {
            assertEquals(expected[i].getLostTime(),actual[i].getLostTime());
        }
    }

    @Test
    void selectByKeyword() {
        LostThing thing1 = new LostThing("重庆");
        LostThing thing2 = new LostThing("成都");
        LostThing thing3 = new LostThing("上海");
        LostThing thing4 = new LostThing("提瓦特大陆");
        LostThing[] expected = new LostThing[]{thing4,thing3,thing2,thing1};

        assertNotNull(s.selectByKeyword(expected,"提瓦特"));
    }
}