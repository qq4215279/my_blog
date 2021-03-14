const body = document.getElementsByTagName("body")[0];

const changeTheme = function(theme){
    let isDark = theme == "dark"
    body.style.setProperty("--backgroundColor", isDark ? "#121212" : "#fff")
    body.style.setProperty("--textColor", isDark ? "#c9d1d9" : "#121212")
    body.style.setProperty("--borderColor", isDark ? "#21262d" : "#f6f6f6")
    //
    body.style.setProperty("--left-menu-bg-color", isDark ? "#121212" : "#20222a")
    body.style.setProperty("--left-border-color", isDark ? "rgba(33,38,45,0.9)" : "rgba(0,0,0,0.15)")
}   
export default changeTheme