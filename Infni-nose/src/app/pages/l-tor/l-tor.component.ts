import { Component } from '@angular/core';
import { NzCardModule } from 'ng-zorro-antd/card';

@Component({
  selector: 'app-l-tor',
  imports: [
    NzCardModule
  ],
  templateUrl: './l-tor.component.html',
  styleUrl: './l-tor.component.less'
})
export class LTorComponent {
  title="Infni Community";
  subtitle="探索无限可能得星球"
}
