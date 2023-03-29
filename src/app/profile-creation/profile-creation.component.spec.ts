import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileCreationComponent } from './profile-creation.component';

describe('ProfileCreationComponent', () => {
  let component: ProfileCreationComponent;
  let fixture: ComponentFixture<ProfileCreationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProfileCreationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProfileCreationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
