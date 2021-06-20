import Parent from "./Parent.js";

export default class CreatePost extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }

 async init() {
    console.log('createPost')
    //use super.fetch(endpoit, object)
    document.querySelector("#createPost").onclick = async () => {
      let post = {
        thread: "",
        content: "",
        category: 0,
      };
      //DRY
      post = super.fillObject(post);
      if (post == null) {
        super.showNotify("post fill error", "error");
        return;
      }
      post.creatorid = document.cookie.split(`; user_id=`).pop().split(';').shift()
      let status = await super.fetch("post/create", post);
      if (status == 200) {
        window.location.replace('/all')
        //redirect -> created post, /post/id
      } else {
        super.showNotify("not create post", "error");
      }
    };
  }

  async getHtml() {
    let body = `
        <div>
        <input id="thread" required placeholder='thread post'/>
        <textarea id="content" required placeholder='content'> </textarea>
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
    let header = super.showHeader("auth");
    return header + body;
  }
}
