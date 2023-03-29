import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NftCardPageComponent } from './nft-card-page.component';

describe('NftCardPageComponent', () => {
  let component: NftCardPageComponent;
  let fixture: ComponentFixture<NftCardPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NftCardPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NftCardPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
