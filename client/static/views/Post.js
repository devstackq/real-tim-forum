import Parent from "./Parent.js";

export default class Post extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }

  1todo, post create
2    use  super.createElement([

  async init() {
    //use super.fetch(endpoit, object)

    document.querySelector("#createPost").onclick = () => {
      let post = {
        thread: "",
        content: "",
        categoryid: 0,
        category: "",
        creatorid: 0,
      };
      //get value from html input
      // post.thread =
      let object = await super.fetch("post/create", post);
      if (object != null) {
        console.log("post created");
      }
    };
  }
  async getHtml() {
    let wrapper = `
        <div>
        <input id="thread" required placeholder='thread post'/>
        <input id="content" required placeholder='content'/>
        <input id="category" required placeholder='thread post'/>
        <button id='createPost'> create </button>
        </div>
        `;
    //authType
    let h = super.showHeader("free");
    return h + wrapper;
  }
}
