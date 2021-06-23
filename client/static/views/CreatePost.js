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
        document.querySelector("#createPost").onclick = async() => {
            let post = {
                thread: "",
                content: "",
                creatorid: 0,
            };
            post = super.fillObject(post);
            post.categories = []

            var inputs = document.querySelectorAll('.category');

            for (var i = 0; i < inputs.length; i++) {
                if (inputs[i].checked) {
                    post.categories.push(inputs[i].value)
                }
            }

            if (post == null) {
                super.showNotify("post fill error", "error");
                return;
            }
            console.log(post)

            let status = await super.fetch("post/create", post);
            // console.log(status, 'create post')
            if (status == 200) {
                window.location.replace('/all')
                    //redirect -> created post, /post/id
            } else {
                console.log(status)
                super.showNotify("session expires or bad request", "error");
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
        </div>
        `;
        let header = super.showHeader("auth");
        return header + body;
    }
}