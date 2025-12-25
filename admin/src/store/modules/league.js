import { getList as fetchLeagueList, submit as submitLeague } from '@/api/league'

/**
 * 获取联赛列表的 Vuex 模块
 * 暴露 state/list、mutations/SET_LIST、actions/getList
 */
export default {
  namespaced: true,
  state: {
    list: []
  },
  mutations: {
    /**
     * 设置联赛列表
     * 入参：state，list 联赛数组
     * 返回：无
     */
    SET_LIST(state, list) {
      state.list = list || []
    }
  },
  actions: {
    /**
     * 获取联赛列表并写入 state
     * 入参：context（含 commit），payload（可选）
     * 返回：Promise，resolve 为联赛数组
     */
    getList({ commit }) {
      return fetchLeagueList().then(res => {
        const data = res.data?.data || res.data || []
        commit('SET_LIST', data)
        return data
      })
    },
    /**
     * 提交联赛数据
     * 入参：context（含 commit），payload（包含联赛数据）
     * 返回：Promise，resolve 为提交结果
     */
    submit({ commit }, data) {
      return submitLeague(data)
    }
  }
}

