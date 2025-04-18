import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LTorComponent } from './l-tor.component';

describe('LTorComponent', () => {
  let component: LTorComponent;
  let fixture: ComponentFixture<LTorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LTorComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LTorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
