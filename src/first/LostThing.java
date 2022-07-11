package first;

import java.math.BigInteger;

public class LostThing {
    private String name;
    private BigInteger lostTime;
    private byte[] photo;
    private String toAddress;

    public LostThing(String name, BigInteger lostTime, byte[] photo, String toAddress) {
        this.name = name;
        this.lostTime = lostTime;
        this.photo = photo;
        this.toAddress = toAddress;
    }

    public LostThing(BigInteger lostTime) {
        this.lostTime = lostTime;
    }

    public LostThing(String toAddress) {
        this.toAddress = toAddress;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public BigInteger getLostTime() {
        return lostTime;
    }

    public void setLostTime(BigInteger lostTime) {
        this.lostTime = lostTime;
    }

    public byte[] getPhoto() {
        return photo;
    }

    public void setPhoto(byte[] photo) {
        this.photo = photo;
    }

    public String getToAddress() {
        return toAddress;
    }

    public void setToAddress(String toAddress) {
        this.toAddress = toAddress;
    }
}
