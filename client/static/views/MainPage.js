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
        console.log(path)
        let response = await fetch(`http://localhost:6969${path}`)
        let result = await response.json()

        console.log(response.statusText, result)

        if (response.status == 200) {
            console.log('main page func, get all posts or category posts show')
        }
    }
    async getHtml() {
        // let uuid = document.cookie.split(";")[1].slice(9, )
        //show list last created post, filter page getAllPost()
        let authState = localStorage.getItem('isAuth')

        if (authState == 'true') {
            // console.log(authState, 'auth state')
            return super.showHeader('auth');
        } else {
            return super.showHeader('free');
        }
    }
}