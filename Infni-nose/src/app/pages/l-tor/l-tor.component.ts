import { Component, inject, OnDestroy, OnInit } from '@angular/core';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzGridModule } from 'ng-zorro-antd/grid';
import {
  NonNullableFormBuilder,
  ReactiveFormsModule,
  Validators
} from '@angular/forms';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzIconModule } from 'ng-zorro-antd/icon';
import {LoginComponent} from '../../components/login/login.component';
import {RegisterComponent} from '../../components/register/register.component';

@Component({
  selector: 'app-l-tor',
  imports: [
    NzCardModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCheckboxModule,
    NzFormModule,
    NzInputModule,
    NzSelectModule,
    NzDividerModule,
    NzIconModule,
    NzGridModule,
    RegisterComponent,
    LoginComponent,
  ],
  templateUrl: './l-tor.component.html',
  styleUrl: './l-tor.component.less'
})
export class LTorComponent {
  title="Infni Community";
  subtitle="探索无限可能得星球"
  isLogin = true;

  private login = inject(NonNullableFormBuilder);
  validateLoginForm = this.login.group({
    email: this.login.control('', [Validators.email, Validators.required]),
    password: this.login.control('', [Validators.required]),
  });


  submitForm(): void {
    if (this.validateLoginForm.valid) {
      console.log('submit', this.validateLoginForm.value);
    } else {
      Object.values(this.validateLoginForm.controls).forEach(control => {
        console.log('control', control);
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  switcher(){
    this.isLogin = !this.isLogin;
  }
}
