import Parent from "./Parent.js";

export default class ViewPost extends Parent {
    constructor(params) {
        super();
        this.params = params;
        this.post = {
            id: window.location.href.split("=")[1],
        }
    }

    setTitle(title) {
        document.title = title;
    }

    async postById() {

            let object = await super.fetch("post/id", this.post);

            let parent = document.querySelector("#postParent");

            if (object != null) {
                let btnTextarea = document.createElement("button");
                btnTextarea.textContent = "lost comment";

                for (let [k, v] of Object.entries(object)) {
                    let span = document.createElement("span");
                    if (v != null) {
                        span.textContent = ` ${k} : ${v} \n`;
                    }
                    parent.append(span);
                }
            } else {
                super.showNotify('bad request', 'error')
                    // parent.innerHTML = ""
            }
        }
        //out -> Parent -> use Comment & Post component
        // async lostComment() {
        //     let content = document.getElementById("commentField").value;
        //     let comment = { content: "", userid: 0, postid: 0 };
        //     comment.content = content;
        //     let object = await super.fetch("comment", comment);
        //     if (object != null) {
        //         //like postById
        //         console.log("lost comment in post, who ?", object);
        //     }
        // }

    async init() {

        let parent = document.querySelector("#postParent");
        // let id  = url.searchParams.get("id")
        this.postById(this.post.id);

        super.createElement([
            { type: "button" },
            { id: "btnlike" },
            { text: "like" },
            { parent: parent },
            { value: 'like' },
            { func: this.postLike },
        ]);

        super.createElement([
            { type: "button" },
            { id: "btndislike" },
            { text: "dislike" },
            { parent: parent },
            { func: this.postDislike },
        ]);
        //init event onclick
        super.setPostParams("post", this.post.id)
        super.voteLike()
        super.voteDislike()
    }

    async getHtml() {
        let body = `<div id='postParent'>  </div>`;
        let comment = `<textarea id="commentField"> </textarea>  <button id="btncomment">lost comment </button>`
        let authState = localStorage.getItem("isAuth");
        if (authState == "true") {
            return super.showHeader("auth") + body + comment
        } else {
            return super.showHeader("free") + body
        }
    }
}