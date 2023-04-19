import { DashboardComponent } from "src/app/components/dashboard/dashboard.component";

describe('dashboard.cy.ts', () => {
  it('playground', () => {
    cy.mount(DashboardComponent);
  })
})