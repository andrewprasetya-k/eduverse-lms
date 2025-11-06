import { Injectable } from '@nestjs/common';
import { SupabaseService } from './supabase/supabase.service';

@Injectable()
export class AppService {
  constructor(private readonly supabase: SupabaseService) {}

  getHello(): string {
    return 'Hello World!';
  }

  // Test connection to Supabase
  async testSupabaseConnection() {
    try {
      const { data, error } = await this.supabase.db
        .from('_test')
        .select('*')
        .limit(1);

      if (error) {
        return {
          status: 'connected',
          message: 'Supabase client initialized (no tables yet)',
          error: error.message,
        };
      }

      return {
        status: 'connected',
        message: 'Supabase connection successful',
      };
    } catch (error) {
      return {
        status: 'error',
        message: error.message,
      };
    }
  }
}
