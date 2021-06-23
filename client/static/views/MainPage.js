import Parent from "./Parent.js";

export default class Posts extends Parent {
    constructor(params) {
        super();
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    render(item, idx, where, type) {
        let wrapper = document.querySelector(where);
        let btn = document.createElement("button");
        let div = document.createElement("div");
        // let className=''
        // let f = null
        //     if(type =='post') {
        //         className = 'postWrapper'
        //         f = this.postById
        //     }else if (type=='profile'){
        //         className = 'profileWrapper'
        //         f = this.editProfile
        div.className = "postWrapper";

        for (let [k, v] of Object.entries(item)) {
            if (v != null) {
                let span = document.createElement("span");
                span.id = k;
                span.textContent = ` ${k} : ${v} `;
                btn.value = idx;
                btn.textContent = `see post`;

                btn.onclick = () => {
                    window.location.replace(`/postget?id=${item["id"]}`)
                        // this.postById(item["id"]);
                };

                div.appendChild(span);
                div.appendChild(btn);
                wrapper.appendChild(div);
            }
        }
    }

    //async - await -> then hell
    async init() {
        let response = await fetch(`http://localhost:6969/api/post/`);
        if (response.status == 200) {
            let result = await response.json();
            //show all post
            result.forEach((element, idx) => {
                this.render(element, idx, ".postContainer");
            });
        }
    }

    async getHtml() {
        // let uuid = document.cookie.split(";")[1].slice(9, )
        let authState = localStorage.getItem("isAuth");
        let body = `<div class="postContainer"> </div>`;

        let create = `<a href="/postcreate" class="nav__link" data-link>create post</a>`
        console.log(authState)
        if (authState == "true") {
            return super.showHeader("auth") + create + body
        } else {
            return super.showHeader("free") + body + '';
        }
    }
}