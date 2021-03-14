const getters = {
    website: state => state.common.website,
    userInfo: state => state.user.userInfo,
    menu: state => state.user.menu,
    topMenu: state => state.user.topMenu,
    tag: state => state.tags.tag,
    tagList: state => state.tags.tagList,
  }
  export default getters