import Parent from "./Parent.js";

export default class Posts extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    init() {

        let category = localStorage.getItem('category')
            //category switch -> route -> requet handler   -> getAllPost, else if /love -> get posts by love    
        console.log('main page func, get all posts or category posts show')
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