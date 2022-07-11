package first;

import java.math.BigInteger;

public class LastBook extends LostThing{
    private String infoForMaster;//书籍可能附带丢失者信息


    public LastBook(String name, BigInteger lostTime, byte[] photo, String toAddress, String infoForMaster) {
        super(name, lostTime, photo, toAddress);
        this.infoForMaster = infoForMaster;
    }

    public LastBook(String name, BigInteger lostTime, byte[] photo, String toAddress) {
        super(name, lostTime, photo, toAddress);
    }

    public String getInfoForMaster() {
        return infoForMaster;
    }

    public void setInfoForMaster(String infoForMaster) {
        this.infoForMaster = infoForMaster;
    }
}
