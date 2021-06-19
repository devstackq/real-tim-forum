import Parent from "./Parent.js";

export default class Post extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title =  title;
  }

  
   init() {
    //use super.fetch(endpoit, object)h
    document.querySelector("#createPost").onclick = async() => {
      let post = {
        thread: "",
        content: "",
        category: 0,
      };
//DRY
post = super.fillObject(post)
if(post == null) { 
 super.showNotify('post fill error', 'error') 
return
}

      // post.creatorid = document.cookie('user_id')
      //get value from html input
      let object = await super.fetch("post/create", post);
      if (object != null) {
        console.log("post created");
        //redirect -> created post
      }else {
        super.showNotify('not create post', 'error')
      }
    };
  }

  async getHtml() {
    let body = `
        <div>
        <input id="thread" required placeholder='thread post'/>
        <input id="content" required placeholder='content'/>
        <label> category:
        <select id='category' required>
        <option></option>
        <option value='2'>love</option>
        <option value='1'>science</option>
        <option value='3'>nature</option>
      </select>
      </label>
        <button id='createPost'> create </button>
        </div>
        `;
    let header = super.showHeader("free");
    return header + body;
  }
}
