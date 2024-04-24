/**
 * 图书类型实体
 */


public class BookType{
	private Integer id;				//类型编号
	private String typeName;

	
	public void setId(Integer id) {
		this.id = id; 
	}

	public void setTypeName(String typeName) {
		this.typeName = typeName; 
	}

	public Integer getId() {
		return (this.id); 
	}

	public String getTypeName() {
		return (this.typeName); 
	}		//类型名
	
	
}