import java.util.*;

/**
 * ����ʵ��
 */


public class Orders{
	private Integer id;				//�������
	private Integer uid;			//�û����
	private String pay;				//���ʽ
	private String carry;			//�ʼķ�ʽ
	private String address;			//�ʼĵ�ַ
	private Date orderDate;			//������Ч����
	private String bz;				//��ע��Ϣ
	private boolean enforce;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setUid(Integer uid) {
		this.uid = uid; 
	}

	public void setPay(String pay) {
		this.pay = pay; 
	}

	public void setCarry(String carry) {
		this.carry = carry; 
	}

	public void setAddress(String address) {
		this.address = address; 
	}

	public void setOrderDate(Date orderDate) {
		this.orderDate = orderDate; 
	}

	public void setBz(String bz) {
		this.bz = bz; 
	}

	public void setEnforce(boolean enforce) {
		this.enforce = enforce; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public Integer getUid() {
		return (this.uid); 
	}

	public String getPay() {
		return (this.pay); 
	}

	public String getCarry() {
		return (this.carry); 
	}

	public String getAddress() {
		return (this.address); 
	}

	public Date getOrderDate() {
		return (this.orderDate); 
	}

	public String getBz() {
		return (this.bz); 
	}

	public boolean getEnforce() {
		return (this.enforce); 
	}		//�Ƿ���ִ��
	
	
	
	
}