function ToggleCockie(name_cockie){
    if( "true" == Cookies.get(name_cockie)){
      Cookies.set(name_cockie, "false", { expires: 365, path: '/' })
      console.log(name_cockie,"OFF")
    }else{
      Cookies.set(name_cockie, "true", { expires: 365, path: '/' })
      console.log(name_cockie,"ON")
    }
    return Cookies.get(name_cockie)
}
