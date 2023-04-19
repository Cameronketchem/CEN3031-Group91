
describe('Connect Wallet Component', () => {
  it('Should create the app', () => {
    cy.visit('http://localhost:4200/connect-wallet')
    cy.get('app-root').should('exist')
  });

  it('User has Metamask extension installed', () => {
    cy.visit('http://localhost:4200/connect-wallet')
    cy.window().then((window) => {
      if(cy.window().ethereum){
        cy.get('.connect-wallet').click()
        cy.wrap(window.send).should('be.called')
      }
    })});

    it('Goes to home page when clicked', () => {
      cy.visit('http://localhost:4200/connect-wallet')
      cy.get('.app-title').click()
      cy.url().should('include', 'http://localhost:4200');
    });

    it('User does not have Metamask extension installed', () =>{
      cy.visit('http://localhost:4200/connect-wallet')
      cy.window().ethereum = undefined
      cy.get('.connect-wallet').click()
      cy.on('window:alert',(popup)=>{
        expect(popup).to.contains('You don\'t have the Metamask extension installed!  Please visit "https://metamask.io/" to create an account!');
     })
    })
})