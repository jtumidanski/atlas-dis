package com.atlas.dis.database.provider;

import com.app.database.util.QueryProviderUtil;
import com.atlas.dis.database.transformer.MonsterDropTransformer;
import com.atlas.dis.entity.MonsterDrop;
import com.atlas.dis.model.MonsterDropData;

import javax.persistence.EntityManager;
import javax.persistence.TypedQuery;
import java.util.List;

public final class MonsterDropProvider {
   private MonsterDropProvider() {
   }

   public static List<MonsterDropData> getAll(EntityManager entityManager) {
      TypedQuery<MonsterDrop> query = entityManager.createQuery("SELECT d FROM MonsterDrop d", MonsterDrop.class);
      return QueryProviderUtil.list(query, new MonsterDropTransformer());
   }

   public static List<MonsterDropData> getByMonsterId(EntityManager entityManager, int monsterId) {
      TypedQuery<MonsterDrop> query = entityManager.createQuery("SELECT d FROM MonsterDrop d WHERE d.monsterId = :monsterId", MonsterDrop.class);
      query.setParameter("monsterId", monsterId);
      return QueryProviderUtil.list(query, new MonsterDropTransformer());
   }
}
