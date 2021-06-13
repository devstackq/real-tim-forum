export default class Parent {
    constructor(text, type) {
        this.text = text
        this.type = type
        this.value = ''
    }

    showHeader(type) {
        console.log(type, 'hedaer')

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
        <a href="/" class="nav__link" data-link>Main</a>
        ${login}
        ${register}
        <div class="dropdown">
          <button class="dropbtn">Categories</button>
          <div class="dropdown-content">
          <a href="#">category 1</a>
          <a href="#">category 2</a>
          <a href="#">category 3</a>
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

    // toggleLinks(type) {
    //     console.log(type, 'toggle')
    //     this.value = type

    //     let signin = document.getElementsByClassName('signin')[0]
    //     let signup = document.getElementsByClassName('signup')[0]
    //     let logout = document.getElementsByClassName('logout')[0]

    //     if (type === 'logout') {
    //         console.log('sginin type, hide links')
    //         signin.style.display = 'block'
    //         signup.style.display = 'block'
    //         logout.style.display = 'none'
    //     } else if (type === 'signin') {
    //         signin.style.display = 'none'
    //         signup.style.display = 'none'
    //         logout.style.display = 'block'
    //     }
    // }
}