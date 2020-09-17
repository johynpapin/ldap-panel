export default store => {
  store.subscribeAction({
    error (action, state, error) {
      store.dispatch('alerts/createAlertFromError', error)
    }
  })
}
