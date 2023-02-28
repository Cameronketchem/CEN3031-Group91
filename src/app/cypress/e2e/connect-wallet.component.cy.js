describe('Tests the Meta mask implementation', () => {
  it('User has Metamask extension installed', () => {
    cy.visit('http://localhost:4200/connect-wallet')
    cy.get('.welcome-msg-background').invoke('css', 'display', 'none');
    cy.window().then((window) => {
      if(window.ethereum){
        cy.get('.connect-wallet').click()
        cy.wrap(window.send).should('be.called')
      }
    })
  })

  it('User does not have Metamask extension installed', () =>{
    cy.visit('http://localhost:4200/connect-wallet')
    cy.get('.welcome-msg-background').invoke('css', 'display', 'none');
    cy.window().ethereum = undefined
    cy.get('.connect-wallet').click()
    cy.on('window:alert',(popup)=>{
      expect(popup).to.contains('You don\'t have the Metamask extension installed!  Please visit "https://metamask.io/" to create an account!');
   })
  })
})