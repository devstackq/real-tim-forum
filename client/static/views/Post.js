import Parent from "./Parent.js";

export default class Post extends Parent {

    constructor(params) {
        super()
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }

    async init() {

        console.log('her')
        let postId = { value: 0 }
        postId.value = window.postId
        console.log(postId, res)
        let response = await fetch('http://localhost:6969/api/post/id', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(postId)
        });
        let res = await response.json(0)

        if (response.status == 200) {
            console.log(response.statusText)
        }
    }

    async getHtml() {
        let wrapper = `
        <div>
        <span id=""> field post</span>
        <span id=""> </span>
        <span id=""> </span>
        <span id=""> </span>
        </div>
        `
            //authType
        let h = super.showHeader('free');
        return h + wrapper
    }
};