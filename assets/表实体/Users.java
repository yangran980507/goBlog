/**
 * 用户实体
 */


public class Users{
	private Integer id;				//用户编号
	private String loginName;		//登录名
	private String trueName;		//真实名
	private String password;		//密码
	private String city;			//所在城市
	private String address;			//所在地址
	private String postcode;		//邮政编码
	private String cardNo;			//证件号码
	private String cardType;		//证件类型
	private Integer grade;			//等级编号(折扣表外键)
	private Float amount;			//消费金额
	private String phone;			//联系电话
	private String email;			//电子邮件
	private boolean freeze;			//帐户是否可用
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
	}			//用户身份(1普通用户，2网站管理员)
	
	
}