/**
 * �û�ʵ��
 */


public class Users{
	private Integer id;				//�û����
	private String loginName;		//��¼��
	private String trueName;		//��ʵ��
	private String password;		//����
	private String city;			//���ڳ���
	private String address;			//���ڵ�ַ
	private String postcode;		//��������
	private String cardNo;			//֤������
	private String cardType;		//֤������
	private Integer grade;			//�ȼ����(�ۿ۱����)
	private Float amount;			//���ѽ��
	private String phone;			//��ϵ�绰
	private String email;			//�����ʼ�
	private boolean freeze;			//�ʻ��Ƿ����
	private Integer status;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setLoginName(String loginName) {
		this.loginName = loginName; 
	}

	public void setTrueName(String trueName) {
		this.trueName = trueName; 
	}

	public void setPassword(String password) {
		this.password = password; 
	}

	public void setCity(String city) {
		this.city = city; 
	}

	public void setAddress(String address) {
		this.address = address; 
	}

	public void setPostcode(String postcode) {
		this.postcode = postcode; 
	}

	public void setCardNo(String cardNo) {
		this.cardNo = cardNo; 
	}

	public void setCardType(String cardType) {
		this.cardType = cardType; 
	}

	public void setGrade(Integer grade) {
		this.grade = grade; 
	}

	public void setAmount(Float amount) {
		this.amount = amount; 
	}

	public void setPhone(String phone) {
		this.phone = phone; 
	}

	public void setEmail(String email) {
		this.email = email; 
	}

	public void setFreeze(boolean freeze) {
		this.freeze = freeze; 
	}

	public void setStatus(Integer status) {
		this.status = status; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public String getLoginName() {
		return (this.loginName); 
	}

	public String getTrueName() {
		return (this.trueName); 
	}

	public String getPassword() {
		return (this.password); 
	}

	public String getCity() {
		return (this.city); 
	}

	public String getAddress() {
		return (this.address); 
	}

	public String getPostcode() {
		return (this.postcode); 
	}

	public String getCardNo() {
		return (this.cardNo); 
	}

	public String getCardType() {
		return (this.cardType); 
	}

	public Integer getGrade() {
		return (this.grade); 
	}

	public Float getAmount() {
		return (this.amount); 
	}

	public String getPhone() {
		return (this.phone); 
	}

	public String getEmail() {
		return (this.email); 
	}

	public boolean getFreeze() {
		return (this.freeze); 
	}

	public Integer getStatus() {
		return (this.status); 
	}			//�û����(1��ͨ�û���2��վ����Ա)
	
	
}