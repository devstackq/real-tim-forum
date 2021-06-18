import Parent from "./Parent.js";

export default class Posts extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }

    async init() {

        let path = window.location.pathname
            // document.cookie( sessionCookie)
        let response = await fetch(`http://localhost:6969/api/posts/`)
            // console.log(`http: //localhost:6969/api/posts/${path}`)
        console.log(response)

        if (response.status == 200) {
            let result = await response.json()
            console.log(result)
            result.forEach((element, idx)  => {
            super.render(element, idx, '.postContainer')
        })
    }
    }
    async getHtml() {
        // let uuid = document.cookie.split(";")[1].slice(9, )
        //show list last created post, filter page getAllPost()
        let authState = localStorage.getItem('isAuth')

        let wrapper = `
        <div class="postContainer">
            
            </div>
            `
        let h = ""
        if (authState == 'true') {
            // console.log(authState, 'auth state')
            h = super.showHeader('auth');
        } else {
            h = super.showHeader('free');
        }
        return h + wrapper
    }
}