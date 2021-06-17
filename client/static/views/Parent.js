export default class Parent {

    constructor(text, type) {
        this.text = text
        this.type = type
        this.value = ''
    }

    render(data, where) {

        let wrapper = document.querySelector(where)

        for (let i = 0; i < wrapper.children.length; i++) {
            correct render
            for (let [k, v] of Object.entries(data)) {
                console.log(k, wrapper.children[i].id)
                if (k == wrapper.children[i].id) {
                    wrapper.children[i].textContent = ` ${k} : ${v}`
                }
            }
        }
        console.log('render', data)
    }



    showHeader(type) {

        let login = ""
        let register = ""
        let logout = ""
        let profile = ""

        if (type == 'free') {
            profile = ""
            logout = ""
            register = `<a href="/signup"  class="nav__link signup" data-link>Signup</a>`
            login = `<a href="/signin"  class="nav__link signin" data-link>Signin</a>`
        } else if (type == 'auth') {
            register = ""
            login = ""
            logout = `<a href="/logout" id='logout' class="nav__link logout" data-link>Logout</a>`
            profile = `<a href="/profile" class="nav__link" data-link>Profile</a>`
        }

        return `
        <nav class="nav">
        <a href="/all" class="nav__link" data-link>Main</a>
        ${login}
        ${register}
        <div class="dropdown">
          <button class="dropbtn">Categories</button>
          <div class="dropdown-content">
          <a href="/love" data-link>love</a>
          <a href="/science" data-link>science</a>
          <a href="/nature" data-link>nature</a>
        </div>
        </div>
       ${profile}
        ${logout}
    </nav>
    
    <span class='notify' > </span> 
`
    }

    showNotify(text, type) {

        let notify = document.getElementsByClassName('notify')[0]

        if (type == 'error') {
            notify.style.display = 'block'
            notify.textContent = text
        } else if (type == 'hide') {
            notify.style.display = 'none'
        }
    }
}