import Parent from "./Parent.js";
import { global } from './Signin.js'

export default class Main extends Parent {
    constructor(params) {
        super()
        this.params = params;
    }

    setTitle(title) {
        document.title = title;
    }
    init() {
        console.log('main page func, get all posts')
    }
    async getHtml() {
        let uuid = document.cookie.split(";")[1].slice(9, )
        console.log(uuid, 'uuid')

        1
        if menu - > signin system ? check uuid == Db Uuid - > show 1 variant
        else another
        2 signin correct - > save global variable - > isAuth = true, when Logout - > isAuth false
        fix button logout - > null ?

            //show list last created post, filter page getAllPost()
            // console.log(global, 'global', window.tesrt)
            console.log(window.isAuth, 'auth state')

        if (window.isAuth) {
            return super.showHeader('auth');
        } else {
            return super.showHeader('free');
        }
    }
}