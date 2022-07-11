package first;

import java.math.BigInteger;

public class LostCard extends LostThing{
    private String cardId;
    private String address;
    private byte[] avatar;

    public LostCard(String name, BigInteger lostTime, byte[] photo, String toAddress, String cardId, String address, byte[] avatar) {
        super(name, lostTime, photo, toAddress);
        this.cardId = cardId;
        this.address = address;
        this.avatar = avatar;
    }

    public String getCardId() {
        return cardId;
    }

    public void setCardId(String cardId) {
        this.cardId = cardId;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public byte[] getAvatar() {
        return avatar;
    }

    public void setAvatar(byte[] avatar) {
        this.avatar = avatar;
    }
}
