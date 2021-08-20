import Parent from "./Parent.js";
import { redirect } from "../index.js";
export default class CreatePost extends Parent {
  constructor(params) {
    super();
    this.params = params;
  }

  setTitle(title) {
    document.title = title;
  }

  async init() {
    document.querySelector("#createPost").onclick = async () => {
      let post = {
        thread: "",
        content: "",
        creatorid: 0,
      };
      post = super.fillObject(post);
      post.categories = [];
      console.log(post);
      var inputs = document.querySelectorAll(".category");

      for (var i = 0; i < inputs.length; i++) {
        if (inputs[i].checked) {
          post.categories.push(inputs[i].value);
        }
      }
      if (post == null) {
        super.showNotify("post fill error", "error");
        return;
      }
      let response = await super.fetch("post/create", post);
      console.log(response, "post");
      if (response == "success") {
        redirect("all");
        // or redirect -> created post, /post/id
      } else {
        let r = await response.json();
        super.showNotify(r.value, "error");
      }
    };
  }

  async getHtml() {
    let body = `
        <div>
        <input id="thread" required placeholder='thread post'/>
        <textarea id="content" required placeholder='content'> </textarea>
        <label> love:
        <input class="category" id="loveId" required  type='checkbox' value="love"/>
        </label>
        <label> science:
        <input class="category" id="scienceId" required  type='checkbox' value="science"/>
        </label>
        <label> nature:
        <input class="category" id="natureId" required  type='checkbox' value="nature"/>
        </label>
      </select>
      </label>
        <button id='createPost'> create </button>
        </div>`;
    let header = super.showHeader();
    return header + body;
  }
}
